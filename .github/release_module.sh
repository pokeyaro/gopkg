#!/bin/bash

#set -ex

PACKAGE_NAME=$1

# Get the version string (extracted from the version file version.go)
function get_version_all_str() {
  dir=$1
  if [ -f "$dir/version.go" ]; then
    all_version=$(cat $dir/version.go | tr -d '\r\n' | cut -d '"' -f 2-2)
    echo $all_version
  else
    echo ""
  fi
}

# Get the next version number
function get_next_version() {
  dir=$1
  all_version=$(get_version_all_str $dir)
  if [ -z "$all_version" ]; then
    new_version="$PACKAGE_NAME/v0.1.0"
  else
    major_version=$(echo $all_version | cut -d '/' -f 2-2 | cut -d 'v' -f 2-2 | cut -d '.' -f 1-1)
    minor_version=$(echo $all_version | cut -d '/' -f 2-2 | cut -d 'v' -f 2-2 | cut -d '.' -f 2-2)
    patch_version=$(echo $all_version | cut -d '/' -f 2-2 | cut -d 'v' -f 2-2 | cut -d '.' -f 3-3)

    if [ -z "$patch_version" ]; then
      patch_version=0
    fi

    new_minor_version=$((minor_version + 1))
    new_version="$PACKAGE_NAME/v$major_version.$new_minor_version.0"
  fi
  echo $new_version
}

# Write the incremented version to the version file
function incr_version_to_file() {
  dir=$1
  new_version=$(get_next_version $dir)
  cat <<EOT > $dir/version.go
package $PACKAGE_NAME

const version = "$new_version"
EOT
}

# Commit and tag the version
function commit_and_tag_version() {
  dir=$1
  version=$2

  git add $dir

  # Prompt for commit message
  echo -n "Enter commit message for version $version: "
  read -r commit_message

  git commit -m "$commit_message"
  git tag -a $version -m "Release $version"
  git push --tag
  git push
}

# Automatically increment the tag
function auto_incr_tag() {
  dir=$1

  # Get relevant information
  old_version=$(get_version_all_str $dir)
  new_version=$(get_next_version $dir)

  # Update file
  incr_version_to_file $dir

  # Commit and tag to git
  commit_and_tag_version $dir $new_version

  echo "package: $PACKAGE_NAME"
  echo "old_version: $old_version"
  echo "new_version: $new_version"
}

# Call the auto increment tag function
auto_incr_tag $1