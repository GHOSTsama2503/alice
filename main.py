#!/bin/python3
if __name__ == "__main__":

    from modules.logger import log
    log.info("Initializing....")

    import lib.env
    import sys

    if len(sys.argv) == 2 and sys.argv[1] == "--test":
        log.info("Starting in test mode...")
        lib.env.PRODUCTION = False

    log.info("Loading language files...")
    import modules.language as language
    language.load()

    from lib.app import bot, set_bot_commands

    log.info("Starting bot...")
    bot.start()

    log.info("Setting bot commands...")
    set_bot_commands()

    log.info("Started successfully.")
    bot.loop.run_forever()
