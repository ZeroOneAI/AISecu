CUR_SCRIPT_DIR=$(dirname "${BASH_SOURCE}")

IMAGE=$(cat "${CUR_SCRIPT_DIR}/image")
VERSION=$(cat "${CUR_SCRIPT_DIR}/version")

cd "${CUR_SCRIPT_DIR}/../../.."
# TODO platform 별로 쪼개기
docker buildx build --platform linux/amd64 -t "${IMAGE}:${VERSION}" -f "${CUR_SCRIPT_DIR}/Dockerfile" .
