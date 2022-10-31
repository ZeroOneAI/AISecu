#! /bin/bash
if [[ -z ${EXECUTE_FILE} ]]; then
  echo "ENV [EXECUTE_FILE] must be defined"
  exit 1
fi

if [[ -z ${PROJECT_DIR} ]]; then
  ORIGIN_PATH=$(pwd)
  cd "$(dirname "${BASH_SOURCE}")/.."
  PROJECT_DIR=$(pwd);
  cd "${ORIGIN_PATH}"
fi

source "${PROJECT_DIR}/build/utils.sh"
source "${PROJECT_DIR}/build/dir_init.sh"

PROJECT_LIST=$(ls "${CMD_DIR}")

if [[ -z ${PROJECT} ]]; then
  PROJECT=${PROJECT_LIST}
fi

for project in ${PROJECT}; do
  if exist_in_arr "${PROJECT_LIST}" "${project}"; then
    BUILD_SCRIPT_PATH="${BUILD_DIR}/project/${project}/${EXECUTE_FILE}"
    if ls "${BUILD_SCRIPT_PATH}" 1> /dev/null 2> /dev/null; then
      if ${BUILD_SCRIPT_PATH}; then
        echo "Success execute ${EXECUTE_FILE} for project [${project}]"
      else
        echo "Something wrong with execute ${EXECUTE_FILE} for project [${project}]"
        exit 1
      fi
    else
      echo "No ${EXECUTE_FILE} file for project [${project}]"
    fi
  else
    echo "Unknown project [${project}]"
  fi
done
