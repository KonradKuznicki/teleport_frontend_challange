# e2e

## Cypress BDD 

### For development (with gui)
```
root/e2e/
$ yarn cypress open
```

### From docker
```
root/e2e/
$ docker run -it -v $PWD:/e2e -w /e2e cypress/included:9.1.1
```

### CI/CD no gui
```
root/e2e/
$ yarn cypress run
```
