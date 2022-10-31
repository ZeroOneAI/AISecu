if [[ -z ${DIR_INIT_SH} ]]; then
  DIR_INIT_SH=imported

  if [[ -z ${PROJECT_DIR} ]]; then
    ORIGIN_PATH=$(pwd)
    cd "$(dirname "${BASH_SOURCE}")/.."
    PROJECT_DIR=$(pwd);
    cd "${ORIGIN_PATH}"
  fi
  BUILD_DIR="${PROJECT_DIR}/build"
  CMD_DIR="${PROJECT_DIR}/cmd"
  PKG_DIR="${PROJECT_DIR}/pkg"
  DEPLOYMENT_DIR="${PROJECT_DIR}/deployments"
fi
