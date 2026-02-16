if [ -z "$1" ]; then
  echo "Usage: $0 parameter is empty!"
  exit 1
fi

SVC_NAME=$CI_PROJECT_NAME
WORKSPACE_DIR="/home/ubuntu/$SVC_NAME/transfer-files/"
IP_SERVER="103.113.79.180"
PATH_DEPLOY=""

command_to_run=$1
case "$command_to_run" in
  deploy_1)
    sudo ssh -p $PORT_SERVER -o StrictHostKeyChecking=no -i /home/ubuntu/.ssh/id_rsa ubuntu@$IP_SERVER "
    if [ ! -d '$WORKSPACE_DIR' ]; then
      echo "Directory Transfer Files does not exist. Creating now."
      mkdir -p $WORKSPACE_DIR
    fi
    "
    ;;
  deploy_2)
    # Multiple Files
    FILES=(
      "$CI_PROJECT_DIR$DOCKER_DIR/docker-compose.yml"
      "$CI_PROJECT_DIR/.env"
      "$CI_PROJECT_DIR/errorcodes.json"
      # Tambahkan lebih banyak file disini...
    )
    printf "%s\n" "${FILES[@]}" | xargs -I {} sudo scp -i /home/ubuntu/.ssh/id_rsa -P $PORT_SERVER {} ubuntu@$IP_SERVER:$TRANSFER_FILES_DIR

    # Multiple Folders
    FOLDERS=(
      "$CI_PROJECT_DIR/src/app/shared/constants/"
      "$CI_PROJECT_DIR/errorcodes/"
      # Tambahkan lebih banyak folder disini...
    )
    for local_folder in "${FOLDERS[@]}"; do
      # sudo scp -P $PORT_SERVER -i /home/ubuntu/.ssh/id_rsa -r $local_folder ubuntu@$IP_SERVER:$TRANSFER_FILES_DIR
      find $local_folder -type f -print0 | xargs -0 -I {} sudo scp -P $PORT_SERVER -i /home/ubuntu/.ssh/id_rsa -r $local_folder ubuntu@$IP_SERVER:$TRANSFER_FILES_DIR
    done
    ;;
  deploy_3)
    sudo ssh -p $PORT_SERVER -i /home/ubuntu/.ssh/id_rsa ubuntu@$IP_SERVER "
    cd $TRANSFER_FILES_DIR &&
    docker compose pull &&
    docker compose down &&
    docker compose up -d --no-build
    "
    ;;
  *)
    echo "Invalid command. Usage: $0 parameter is empty!"
    exit 1
    ;;
esac
