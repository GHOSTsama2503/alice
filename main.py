#!/bin/python3
from logging import basicConfig, getLogger, log, DEBUG, INFO, WARNING, ERROR, CRITICAL
basicConfig(format="[%(levelname)s]: %(message)s", level=INFO, force=True)
log(INFO, "Initializing...")

from pyrogram import Client
from pyrogram.enums.parse_mode import ParseMode

pyrogram_logger = getLogger("pyrogram")
pyrogram_logger.setLevel(WARNING)

import env
client = Client(
    "bot",
    api_id = env.API_ID,
    api_hash = env.API_HASH,
    bot_token = env.BOT_TOKEN,
)

# Completely disable parsing fo peace of mind.
client.set_parse_mode(ParseMode.DISABLED)

log(INFO, "Loading plugins...")
import plugins_loader
import plugins
plugins_loader.add_plugins(client, plugins)

log(INFO, "Starting...")
client.start()

log(INFO, "Started successfully.")
client.loop.run_forever()
