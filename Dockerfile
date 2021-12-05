FROM node AS frontend

WORKDIR /frontend
COPY frontend/package.json frontend/yarn.lock ./
RUN yarn install --frozen-lockfile
COPY frontend /frontend
RUN yarn build-all


FROM golang AS backend

WORKDIR /backend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /backend/app



FROM scratch

WORKDIR /soluiton
COPY --from=frontend /frontend/bu ./index.js
COPY --from=backend /backend/app ./
ENV NODE_SERVER_PATH="/index.js"
ENTRYPOINT ["./app"]
CMD ["help"]

