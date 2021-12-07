FROM node:14 AS frontend

WORKDIR /frontend
COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install --frozen-lockfile
COPY frontend /frontend
RUN PUBLIC_URL="/files" yarn build
RUN PUBLIC_URL="/login" yarn build-login


FROM golang AS backend

WORKDIR /backend
COPY backend .
RUN go build -o /backend/app



FROM ubuntu

RUN groupadd -r chellenger && useradd -r -g chellenger chellenger
WORKDIR /challenge/backend
COPY resources /challenge/resources
COPY --from=frontend /frontend/build /challenge/frontend/build
COPY --from=frontend /frontend/login-build /challenge/frontend/login-build
COPY --from=backend /backend/app ./
RUN chown -R chellenger:chellenger /challenge/resources/traversable
USER chellenger
ENTRYPOINT ["./app"]


