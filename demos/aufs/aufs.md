```$xslt
mkdir aufs-demo && cd aufs-demo

mkdir container-layer
echo "I am container layer" > container-layer/container-layer.txt

for i in 1 2 3 4
do
	echo $i
	mkdir image-layer-${i}
	echo "I am image layer ${i}" > image-layer-${i}/image-layer-${i}.txt
done

mount -t aufs -o dirs=./container-layer:./image-layer4:./image-layer3:./image-layer2:./image-layer1 none ./mnt
```