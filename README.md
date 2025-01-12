# k8s-static-server

## Installation

```bash
git clone https://github.com/GoWebProd/k8s-static-server.git
helm install -n static --create-namespace static-server ./k8s-static-server
```

## Creating static server

```yaml
apiVersion: static-server.webprod.io/v1
kind: StaticRepo
metadata:
  name: roubi
  namespace: static
spec:
  hostname: roubi.fake.com
  git:
    repo: SOME_REPO
    commit: SOME_COMMIT

---

apiVersion: static-server.webprod.io/v1
kind: StaticRepo
metadata:
  name: voyah
  namespace: static
spec:
  hostname: voyah.fake.com
  cmd:
    image: alpine/git:latest
    command:
      - /bin/sh
      - -c
      - echo $DATA_DIR && mkdir -p $DATA_DIR/cunbaemul/otan/ && echo "999999999" > $DATA_DIR/cunbaemul/otan/activate3

---

apiVersion: static-server.webprod.io/v1
kind: StaticRepo
metadata:
  name: wiki
  namespace: static
spec:
  hostname: docmost.fake.com
  imageDir:
    image: docmost/docmost:0.6.2
    path: /app/apps/client/dist/
```

## Creating ingress

```yaml
kind: Service
apiVersion: v1
metadata:
  name: static-server
  namespace: application
spec:
  type: ExternalName
  externalName: static-server.static.svc.cluster.local

---

apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: static-server
  namespace: application
spec:
  rules:
    - host: roubi.fake.com
      http: &httpTemplate
        paths:
          - pathType: Prefix
            path: /
            backend:
              service:
                name: static-server
                port:
                  number: 80
    - host: voyah.fake.com
      http: *httpTemplate
    - host: docmost.fake.com
      http: *httpTemplate
  tls:
    - hosts:
        - roubi.fake.com
        - voyah.fake.com
        - docmost.fake.com
      secretName: static-server-cert

```
