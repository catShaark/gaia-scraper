# gaia-scraper
script to scavenge all the gov votes of gaia using raw data
# here is how I do it
1. Download the data, config, binary for each chain in this links : [gaia2] (https://archive.interchain.io/hub2.html) , [gaia3] (https://archive.interchain.io/hub3.html) , [gaia1] (https://archive.interchain.io/hub1.html) 
2. Set gaia1 rpc port to 2001, gaia2 rpc port to 2002, gaia3 rpc port to 2003, Run gaia1 in ~/gaia1/.gaiad, run gaia2 in ~/gaia2/.gaiad, run gaia3 in ~/gaia3/.gaiad
3. Query all txs by iterating through every blocks in each of the 3 chains, output to 3 files data1.json, data2.json, data3.json ( you can do so by running scrape.sh)
4. Run main.go to extract all the vote txs among the txs in data1.json, data2.json, data3.json, output to 3 csv files
5. Profit