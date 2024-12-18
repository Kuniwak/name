# from https://github.com/shogo82148/go-mecab/blob/main/.github/install-mecab-mingw.sh
# from https://gist.github.com/dtan4/351d031bec0c3d45cd8f
# see also http://qiita.com/dtan4/items/c6a087666296fbd5fffb

set -euxo pipefail

install-mecab() {
  local prefix="$1"
  local mecab_version='0.996.10'
  local ipadic_version='2.7.0-20070801'

  local tmp_dir
  tmp_dir="$(mktemp -d)"

  export PATH="$prefix/bin:$PATH"

  if ! [[ -e "$prefix/bin/mecab" ]]; then
    # install mecab
    (cd "$tmp_dir"
      curl -o mecab.tar.gz -sSL "https://github.com/shogo82148/mecab/releases/download/v$mecab_version/mecab-$mecab_version.tar.gz"
      tar zxfv mecab.tar.gz
    )

    (cd "$tmp_dir/mecab-$mecab_version"
      ./configure --enable-utf8-only --host=x86_64-w64-mingw32 --prefix="$prefix"
      make -j
      # make check # it fails :(
      make install
    )

    (cd "$tmp_dir"
      curl -o mecab-ipadic.tar.gz -sSL "https://github.com/shogo82148/mecab/releases/download/v$mecab_version/mecab-ipadic-$ipadic_version.tar.gz"
      tar zxfv mecab-ipadic.tar.gz
    )

    (cd "$tmp_dir/mecab-ipadic-$ipadic_version"
      ./configure --with-charset=utf8 --prefix="$prefix"
      make
      make install
    )
  fi

  set-env 'CGO_LDFLAGS' "\"-L$(cygpath -w /mingw64/lib)\" \"-L$(cygpath -w "$prefix/lib")\" -lmecab -lstdc++"
  set-env 'CGO_CFLAGS' "\"-I$(cygpath -w /mingw64/include)\" \"-I$(cygpath -w "$prefix/include")\""

  if ! grep -sF dicdir "$prefix/etc/mecabrc" >/dev/null; then
    cat << DIC > "$prefix/etc/mecabrc"
dicdir = $(cygpath -w "$prefix/lib/mecab/dic/ipadic")
DIC
  fi

  # The default mecabrc path is "C:\Program Files\mecab\etc\mecabrc" if mecab is built with mingw32-w64.
  # but it is not correct in MSYS2 environment.
  set-env 'MECABRC' "$(cygpath -w "$prefix/etc/mecabrc")"
  add-path "$(cygpath -w "/mingw64/bin")"
  add-path "$(cygpath -w "$prefix/bin")"

  # Could not launch mecab-config on Windows because it is a shell script.
  if [[ -e "$prefix/bin/mecab-config" ]]; then
    rm "$prefix/bin/mecab-config" 
  fi
}
