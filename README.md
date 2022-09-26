# Alice

Multi-functional Telegram bot, with a very practical `plugin` system, the `plugins` are nothing more than pyrogram handlers, which makes the insertion process quite simple.

```py
from pyrogram import Client, filters
from pyrogram.types import Message
from pyrogram.handlers import MessageHandler

async def example_handler(client: Client, message: Message, t=MessageHandler, f=filters.command("start")):
    await message.reply("Hello!")
```

The code snippet, in a .py, inside the `./plugins` folder, would be enough to add an event that responds with "Hello!" when sending the `/start` command.

Recommendations, improvements, and insertion of new plugins in the project are accepted.  Please note that it is still in development and there may be big changes.

### Requirements

`api hash` and `api id`: to get them go to [my.telegram.org](https://my.telegram.org/apps) and register a new application

`bot token`: create a bot in [BotFather](https://t.me/BotFather)

### Running steps

Clone the repository:
```sh
git clone https://github.com/GHOSTsama2503/Alice
```

Change to the project folder:
```sh
cd Alice
```

Create virtual environment:
```sh
python3 -m venv .venv
```

Activate the venv:
```sh
source ./.venv/bin/activate
```

Install dependencies:
```sh
pip install -r requirements.txt
```

Put your credentials in an .env file at the root of the project, click [here](https://github.com/GHOSTsama2503/Alice/blob/main/.env.example) to see an example.

And finally:
```sh
python3 main.py
```
