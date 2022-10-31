CUR_SCRIPT_DIR=$(dirname "${BASH_SOURCE}")

IMAGE=$(cat "${CUR_SCRIPT_DIR}/image")
VERSION=$(cat "${CUR_SCRIPT_DIR}/version")

docker push "${IMAGE}:${VERSION}"
