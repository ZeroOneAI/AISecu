if [[ -z ${UTILS_SH} ]]; then
  UTILS_SH=imported

  function exist_in_arr() {
    local ARR_STR=$1
    local ELEM=$2

    for i in ${ARR_STR}; do
      if [[ $ELEM == $i ]]; then
        return 0
      fi
    done
    return 1
  }
fi
