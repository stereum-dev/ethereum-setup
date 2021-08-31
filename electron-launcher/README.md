# Electron Launcher

## Project setup
```
npm install -g modernizr
npm install iconv-corefoundation
npm install
```

### Build electron
```
# ensure that we have the latest libs. If we skip this step it may happen that libs are missing in the final bundle
npm install
# start electron build
npm run electron:build
```
### Test electron
```
npm run electron:serve
```
