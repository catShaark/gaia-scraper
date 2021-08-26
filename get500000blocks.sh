
start=$(($1*500000))
end=$(($1*500000+500000))
file=~/gaia$2/chunk_of_blocks_$1.json
for (( i=$start; i < $end; i ++ ))
do
   curl "localhost:200$2/block?height=$i" | jq '.result.block.data.txs' >> $file
   echo $i >> $file
done

