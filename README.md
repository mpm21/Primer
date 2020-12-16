# MyCookBook
Dockerfile -> go linter
docker build -t <image_name>:<tag> .
	
docker run --rm -v $(pwd):/app -w /app <image_name>:<tag>
