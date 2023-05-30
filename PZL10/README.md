# Arweave Puzzle #10 - "...execture the order 66..."

Here i am going to publish ideas i consider about and dummy instruments related to Arweave puzzle #10.

This puzzle is one from the series of Arweave Puzzles made by @Tiamat (official Arweave project discord member, apps developer, miner, early investor). Originally, puzzle posted in official @Tiamat twitter [[1]](
https://twitter.com/ArweaveP/status/1250025710133092353): with hints and link to the blockchain [[2]](https://d5zw4kksq5gasg7ezkjvdpey562svcviatdltuyga43lrkexvngq.arweave.net/H3NuKVKHTAkb5MqTUbyY77UqiqgExrnTBgc2uKiXq00
)

Puzzle consists of the HTML page (sources also attached in this git - but could be found on Arweave "blockchain") page with an AES encrypted message (the Arweave wallet with **500 AR** coins on it, that is around of **5 k$** right now). The decryption mechanism is already built in page so the solver should only type the correct symbols to the long placeholder as show on picture below. The key to the puzzle solving is an image of some shread picture that likely to show separated pictures\objects:

![AR puzzle #10 keys](https://github.com/HomelessPhD/AR_Puzzles/blob/8bf62655e36fb91f9db36fa8cfe44059f149495a/PZL10/pics/PZL_10.png )

![AR puzzle #10 solving](https://github.com/HomelessPhD/AR_Puzzles/blob/8bf62655e36fb91f9db36fa8cfe44059f149495a/PZL10/pics/PZL_10_solving.png )

# Assembling the puzzle
The puzzle originally given by shread or dissassebled picture - here what one can obtained assembling it back to the one picture

![AR puzzle #10 keys ASSEMBLED](https://github.com/HomelessPhD/AR_Puzzles/blob/f9fa2f2795d35a2183ff00159a8692bb2e17125d/PZL10/pics/pzl_10_assembled.jpg)

# The official hints given by the puzzle creator are listed below:

    
```
    From Twitter:
    

```    

Some ideas came to my mind, others were captured from Twitter or\and the telegram group [[3]](@arweavep) - you definitely should visit it for ideas sharing.
(at first few months i have forcefully avoided visiting any forums or groups when i saw the puzzle at first because i thought that ideas of others could freeze my own thinking process - it is when you start to think primarly in the way proposed by some other solver and, got stuck on it, but later trying my own ideas with my own dummy brute forcing method i skipped that rule. I am actively posting any ideas came to my mind in that group, hoping someone of us will be able to solve the puzzle sharing some part of revenue and GLORY)

There are pretty "classic" ideas of what each image could code, but also there apretty exotic ideas. The last will be discussed at the end of this paper. Lets consider each picture "objects" one by one at first:

## 1-st key

![AR puzzle #10 key #1](https://github.com/HomelessPhD/AR_Puzzles/blob/42008e05471a8821fdf8d1277455f597c14ab6f5/PZL10/pics/pzl10_key_1.jpg)
....

## 2-nd key

![AR puzzle #10 key #2](https://github.com/HomelessPhD/AR_Puzzles/blob/42008e05471a8821fdf8d1277455f597c14ab6f5/PZL10/pics/pzl10_key_2.jpg)
....

## 3-rd key

![AR puzzle #10 key #3](https://github.com/HomelessPhD/AR_Puzzles/blob/42008e05471a8821fdf8d1277455f597c14ab6f5/PZL10/pics/pzl10_key_3.jpg)
....

## 4-th key

![AR puzzle #10 key #4](https://github.com/HomelessPhD/AR_Puzzles/blob/42008e05471a8821fdf8d1277455f597c14ab6f5/PZL10/pics/pzl10_key_4.jpg)
....

## 5-th key

![AR puzzle #10 key #1](https://github.com/HomelessPhD/AR_Puzzles/blob/42008e05471a8821fdf8d1277455f597c14ab6f5/PZL10/pics/pzl10_key_5.jpg)
....


# "Exotic" ideas

1] ....



`to be continued`

# Brute-Force - "that's impossible (?!)" [12](https://www.youtube.com/watch?v=6ixvpLCdqkA)
![Its necessary](https://raw.githubusercontent.com/HomelessPhD/AR_Puzzles/main/PZL3/pics/Cooper_necessary.gif)
![brute it TARS!!!](https://raw.githubusercontent.com/HomelessPhD/AR_Puzzles/main/PZL3/pics/interstellar-cooper.gif)

Tiamat (the puzzle creator) has defended the puzzle from blind bruteforce. The wallet JSON string (that contains the private key) AES-encrypted by 512 bit pass-key. The 512 bit pass-key is a result of 11512 sequential SHA512 evaluations of the "keys" typed by the user - better look page sources or code snippet below.

```
....
function decodewallet(t,e)
{
    for(var i=CryptoJS.SHA512(e),n=0;n<11512;n++)
	    i=CryptoJS.SHA512(i);
		
	CryptoJS.algo.AES.keySize=32,CryptoJS.algo.EvpKDF.cfg.iterations=1e4,CryptoJS.algo.EvpKDF.cfg.keySize=32;
	var r = CryptoJS.AES.decrypt(t,i.toString());
	
	return out=hex2a(r),out
}
....

var msg="U2FsdGVkX1+E2/9....<LOTS OF SYMBOLS - SEE THE SOURCES>"

function proceed() {
	var x = document.getElementsByClassName("inputs");
        var code="";
        for (var i=0;i<x.length;i++)
                code+=x[i].value;

	a=decodewallet(msg,code);
	if (a.search('"kty":"RSA"')>-1) {
		document.getElementById('dstatus').innerHTML="SUCCESS";
		download("arweave_keyfile_G2BaxD9phYHJ55VaEY-aX28FtQCKLORMMQSc74IaqYg.json",a);
	}
	else document.getElementById('dstatus').innerHTML="FAILED";
}

```
First, the bruteforce of the inputs is slown down by 11512 SHA512 + AES decryption. On a typical CPU you could expect around 0.5 sec for one thread (or CPU core) to compute this decrypting in browser (javascript that is in the original page) 

The Second problem - javascript library CryptoJS used here for computing AES decrypting: seems like it has a bug\feature that makes its results unique for some cases. Thus, using 512-bit key the CryptoJS AES gives the result that the typical AES library - not [[13]](https://github.com/brix/crypto-js/issues/293) and so to build some efficient brute force instrument, the coder will need at least to adjust the AES library he use to follow CryptoJS ruined logic. At least, that is how i see it now.

In general, the list of answers to try on this puzzle could be all possible combinations of all allowed symbols that will result in a huge list (assume all numbers {0-9}, letters{a-z}, special chars {,.\|+=...} ~ 50 different symbols or so) == `50^32 ~ 10^54` but that unreal to be brutted even if some efficient code will brute this task on all bitcoin ASICS on the plannet with their typical hashing speed.

Usually, the solver who wants to try to apply brute force here will compose a dictionary of possible "keys" - inputs. All "keys" or inputs mentioned in this README and probably all 4-letter words and abbreviations or terms mentioned on Arweave materials (including twitter, telegram, discord, promo materials, ..., Britain enyclopedia, whatever) - see extramaterials attached in this git. Typically, such a dictionary will be still huge for bruteforce on reasonable resource with the code we have - some of a core telegram members stated to try trillions of combinations on their PCs using some effective code (cant confirm the code efficience and\or correctness). And so, words for brute should be filtered. 

The ideal bruteforce program, i guess, will consist of two parts: the first will compose a list of inputs and - precompute the 11512 SHA512 hashing of the inputs; the second will compute AES with the resulted 512 bit keys and compare the result with "kty":"RSA" substring. The work could be done in chunks: portion of the dictionary transformed into 512-key and transfered for AES evaluation - to resolve problem of memory and probably to paralelize the work in a more efficient manner. SHA512 is easy part while AES is a problem due to mentioned "bug\feature" of CryptoJS lib.

I have composed a simple JavaScript code to do a bruteforce just in browser - thus i could avoid of re-coding CryptoJS features and yet this solution is a snail (very slow - 0.5 sec per core on 1 input). Later i have speed up this approach precomputing sha512 key on the c-program (CPU) and it increased the speed but yet not enough to call this a WIN. JavaScript just killing all mood here. 

Anyway, with my JS approach i bruted a million inputs per day - and no success.

`TO BE WRITTEN VERY SOON (day or two i will fill this section)`


## P.S.

Thank you for spending time on my notes, i hope it was not totally useless and you've found something interesting. 

Any ideas\questions or propositions you may send to generalizatorSUB@gmail.com - also look at my twitter [[11]](https://twitter.com/miningpredict) @MiningPredict.

-------------------------------------------------------------------------
### References:

[1] Original @ArweaveP (@Tiamat) tweet - 
https://twitter.com/ArweaveP/status/1250025710133092353

[2] Arweave Puzzle #10 stored in Arweave "blockchain" - https://2xzm6mh75smp5ivf2img3biam3iu7qhodsw5mk7cimkx7g4trl4a.arweave.net/1fLPMP_smP6ipdIYbYUAZtFPwO4crdYr4kMVf5uTivg

[3] Telegram group of Arweave puzzles solvers community - @arweavep



[11] MiningPredict (my twitter page) - https://twitter.com/miningpredict

[12] "that's impossible (?!) no it's necessary"  - https://www.youtube.com/watch?v=6ixvpLCdqkA

[13] CryptoJS bug - https://github.com/brix/crypto-js/issues/293

-------------------------------------------------------------------------
### Support
I am poor Ukrainian student that will really appreciate any donations.
I have no home (flat\appartment), live in the dorm (refugee shelter).
 
P.S. Successfully evacuated from occupied regions of Ukraine.

**BTC**:  `1QKjnfVsTT1KXzHgAFUbTy3QbJ2Hgy96WU`

**LTC**:  `LNQopZ7ozXPQtWpCPrS4mGGYRaE8iaj3BE`

**DOGE**: `DQvfzvVyb4tnBpkd3DRUfbwJjgPSjadDTb`

**AR**: `0UM6uoLrrnxXuYpHMBDAv-6txNTMdaEkR2m_bP_1HyE`
(have never used Arweave wallet)