num_of_block=$(eval "echo \${height_of_gaia$1}")
num_of_chunk=$(($num_of_block / 500000))
r=$(($num_of_block % 500000))
cmd=""
for (( i = 0; i < $num_of_chunk; i ++ ))
do 
	cmd="${cmd} . ./get500000blocks.sh ${i} $1 &"
done 
cmd="$cmd . ./get_blocks_in_range.sh $(($num_of_block-$r)) $num_of_block $1 &"
eval $cmd

