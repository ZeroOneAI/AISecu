CUR_SCRIPT_DIR=$(dirname "${BASH_SOURCE}")

export IMAGE=$(cat "${CUR_SCRIPT_DIR}/image")
export VERSION=$(cat "${CUR_SCRIPT_DIR}/version")

if [[ -z ${PROJECT_DIR} ]]; then
  ORIGIN_PATH=$(pwd)
  cd $(dirname "${BASH_SOURCE}")/../../..
  PROJECT_DIR=$(pwd);
  cd "${ORIGIN_PATH}"
fi

cat "${PROJECT_DIR}/deployments/kubernetes/alertmanager.yaml" | envsubst | kubectl apply -f -
