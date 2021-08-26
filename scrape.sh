height_of_gaia1=500041
height_of_gaia2=2902001
height_of_gaia3=5200790

for i in {1..3}
do
    
	. ./get_data_of_gaia.sh $i 
done

for x in {1..3}
do
	for i in $(find ~/gaia${x} -type f -name "chunk_of_blocks_*")
	do 
	  echo $i
	  perl -p -e 's/[0-9]+\n//'  $i >> ~/gaia-scraper/data${x}.json
	done
	sed -i 's/null//g' ~/gaia-scraper/data${x}.json
    sed -i ':a;N;$!ba;s/\n//g' ~/gaia-scraper/data${x}.json
	sed -i 's/\]\[/,/g' ~/gaia-scraper/data${x}.json
done