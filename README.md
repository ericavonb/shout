# SHOUT!
Print large block letters. Because when you are frustrated with Go, you need your terminal output to really scream.

## How to use
To use,
`` import "https://github.com/ericavonb/shout"``

To make a string into block letters, use 'Blockify':
``Blockify("BOO!", 80)``
returns:
``"______       ____       ____     _    \n 
|\ _____\   /\ ____\   /\ ____\  |\_\   \n
| |  __  | | /  ___  \| /  ___  \| | |  \n
| | |_\| / | | | | | || | | | | || | |  \n
| |  __  \ | | |_| | || | |_| | | \|_|  \n
| | |_\| | \ | |__\| |\ | |__\| ||\_\   \n
 \|______/   \______/   \______/  \|_|  \n"`
 
 Blockify takes a string and the maximum width.
 
 To print out your message, use 'PrintBlocks':
 
 ``PrintBlocks("No go!") ``
 prints:
 ``__    __     ____             ______       ____     _   \n  
|\ _\ |\ \  /\ ____\          |\ _____\   /\ ____\  |\_\   \n
| |  \| | || /  ___  \        | |  ____| | /  ___  \| | |  \n
| |   \ | || | | | | |        | | |  __  | | | | | || | |  \n
| | |\ \| || | |_| | |        | | |_|\_\ | | |_| | | \|_|  \n
| | | \   |\ | |__\| |        | | |__\| |\ | |__\| ||\_\   \n
 \|_|  \__|  \______/          \|______/   \______/  \|_|  \n``

'PrintBlocks' uses the width of your terminal as the maximum width.

The currently supported characters are all letters , ".", "-", "_", "?", "!"

## Built in functions
Some built-in functions:
- 'MegaFail(err)` prints "FAIL!" if err is not nil.
- 'BadGo()` prints "BAD GO!"


## Coming soon:
Numbers, more symbols, underlining, colors!	
	
