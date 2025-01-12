/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"os"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	staticserverv1 "webprod.io/static-server/api/v1"
)

// StaticRepoReconciler reconciles a StaticRepo object
type StaticRepoReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=static-server.webprod.io,resources=staticrepoes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=static-server.webprod.io,resources=staticrepoes/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=static-server.webprod.io,resources=staticrepoes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the StaticRepo object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.4/pkg/reconcile
func (r *StaticRepoReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch all StaticRepo resources
	var staticRepoList staticserverv1.StaticRepoList
	if err := r.List(ctx, &staticRepoList); err != nil {
		log.Error(err, "unable to list StaticRepo resources")
		return ctrl.Result{}, err
	}

	// Update Deployment and ConfigMap based on the list of StaticRepo resources
	if err := r.updateStaticServer(ctx, staticRepoList.Items); err != nil {
		log.Error(err, "unable to update static server")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *StaticRepoReconciler) updateStaticServer(ctx context.Context, staticRepos []staticserverv1.StaticRepo) error {
	log := log.FromContext(ctx)
	namespace := os.Getenv("WATCH_NAMESPACE")

	// Fetch the existing Deployment
	var deployment appsv1.Deployment
	if err := r.Get(ctx, client.ObjectKey{Namespace: namespace, Name: "static-server"}, &deployment); err != nil {
		log.Error(err, "unable to fetch Deployment")
		return err
	}

	deployment.Spec.Template.Spec.Volumes = deployment.Spec.Template.Spec.Volumes[:1]
	deployment.Spec.Template.Spec.Containers[0].VolumeMounts = deployment.Spec.Template.Spec.Containers[0].VolumeMounts[:1]
	deployment.Spec.Template.Spec.InitContainers = deployment.Spec.Template.Spec.InitContainers[:0]

	for _, repo := range staticRepos {
		volumeName := strings.Replace(fmt.Sprintf("repo-%s", repo.Spec.Hostname), ".", "-", -1)
		initContainer := strings.Replace(fmt.Sprintf("init-%s", repo.Spec.Hostname), ".", "-", -1)

		deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, corev1.Volume{
			Name: volumeName,
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		})

		deployment.Spec.Template.Spec.Containers[0].VolumeMounts = append(deployment.Spec.Template.Spec.Containers[0].VolumeMounts, corev1.VolumeMount{
			Name:      volumeName,
			MountPath: fmt.Sprintf("/usr/share/nginx/html/%s", repo.Spec.Hostname),
		})

		var (
			image   string
			env     []corev1.EnvVar
			command []string
		)

		switch {
		case repo.Spec.Git != nil:
			image = "alpine/git:latest"
			env = []corev1.EnvVar{
				{
					Name:  "REPO",
					Value: repo.Spec.Git.Repo,
				},
				{
					Name:  "COMMIT",
					Value: repo.Spec.Git.Commit,
				},
				{
					Name:  "DATA_DIR",
					Value: "/data",
				},
			}
			command = []string{"/bin/sh", "-c", "git clone $REPO o && cd o && git checkout $COMMIT && cp -r build/* $DATA_DIR && cd .. && rm -rf o"}
		case repo.Spec.Cmd != nil:
			image = repo.Spec.Cmd.Image
			env = []corev1.EnvVar{
				{
					Name:  "DATA_DIR",
					Value: "/data",
				},
			}
			command = repo.Spec.Cmd.Command
		case repo.Spec.ImageDir != nil:
			if repo.Spec.ImageDir.Path[len(repo.Spec.ImageDir.Path)-1] == '/' {
				repo.Spec.ImageDir.Path = repo.Spec.ImageDir.Path[:len(repo.Spec.ImageDir.Path)-1]
			}

			image = repo.Spec.ImageDir.Image
			env = []corev1.EnvVar{
				{
					Name:  "DATA_DIR",
					Value: "/data",
				},
			}
			command = []string{"/bin/sh", "-c", fmt.Sprintf("cp -r %s/* $DATA_DIR", repo.Spec.ImageDir.Path)}
		}

		deployment.Spec.Template.Spec.InitContainers = append(deployment.Spec.Template.Spec.InitContainers, corev1.Container{
			Name:    initContainer,
			Image:   image,
			Env:     env,
			Command: command,
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      volumeName,
					MountPath: "/data",
				},
			},
		})
	}

	// Update the Deployment
	if err := r.Update(ctx, &deployment); err != nil {
		log.Error(err, "unable to update Deployment")
		return err
	}

	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *StaticRepoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&staticserverv1.StaticRepo{}).
		Named("staticrepo").
		Complete(r)
}
