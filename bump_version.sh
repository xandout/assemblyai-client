#!/usr/bin/env bash
if [[ $1 == '-h' || $1 == '--help' ]];then
  echo 'Create a git tag for the version in VERSION and push. Fail if the version already exists on remote.'
  exit 0
fi
set -euo pipefail
IFS=$'\n\t'
# make sure we're in the correct directory
cd "$(dirname "$0")"
# assign version from VERSION file at project root
VERSION=$(head -n 1 VERSION)
if [[ ! "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+.*$ ]]; then
  echo "version '$VERSION' does not match expected format, inspect VERSION file"
  exit 1
fi
# fail if $VERSION already exists on remote
if git ls-remote --quiet --tags --exit-code origin $VERSION &> /dev/null;then
  echo "tag $VERSION already exists on remote. only unique tags are allowed, please update VERSION file."
  exit 1
fi
# tag current commit
git tag $VERSION
# push new tag to remote
git push origin $VERSION