REPO=ghcr.io
IMAGE=tienvnz98/togo
VERSION=$(cat package.json | grep version | head -1 | awk -F= "{ print $2 }" | sed 's/[version:,\",]//g' | tr -d '[[:space:]]')
IMAGE_NAME=$REPO/$IMAGE

# Pre require login to registry first
# docker login ghrc.io

docker build -t $IMAGE_NAME:$VERSION .
# TODO: Push this image to registry
docker build -t $IMAGE_NAME:latest .

docker push $IMAGE_NAME:$VERSION
docker push $IMAGE_NAME:latest