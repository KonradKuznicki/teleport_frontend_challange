FROM node:14 AS frontend

WORKDIR /frontend
COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install --frozen-lockfile
COPY frontend /frontend
RUN PUBLIC_URL="/files" yarn build
RUN PUBLIC_URL="/login" yarn build-login


FROM golang:1.17 AS backend

WORKDIR /backend
COPY backend .
RUN CGO_ENABLED=0 go build -o /backend/app
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group


FROM scratch

WORKDIR /challenge/backend
COPY --from=backend /user/group /user/passwd /etc/
COPY resources /challenge/resources
COPY --chown=nobody:nobody resources/traversable /challenge/resources/traversable
COPY --from=frontend /frontend/build /challenge/frontend/build
COPY --from=frontend /frontend/login-build /challenge/frontend/login-build
COPY --from=backend /backend/app ./

USER nobody:nobody
ENTRYPOINT ["./app"]


