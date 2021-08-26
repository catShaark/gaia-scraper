n=$(($1/500000))
file=~/gaia$3/chunk_of_blocks_${n}.json
for (( i=$1; i <= $2; i ++ ))
do
   curl "localhost:201$3/block?height=$i" | jq '.result.block.data.txs' >> $file
   echo $i >> $file
done

