import logging
import sys
from datetime import datetime
from pathlib import Path

from termcolor import colored

import app.env as env

DEBUG = env.DEBUG
LOGS_DIR = "logs"


class CustomFormatter(logging.Formatter):
    FORMATS = {
        logging.DEBUG: "%(levelname)s",
        logging.INFO: colored("%(levelname)s", "blue"),
        logging.WARN: colored("%(levelname)s", "yellow"),
        logging.ERROR: colored("%(levelname)s", "red"),
        logging.CRITICAL: colored("%(levelname)s", "red", None, ["bold"]),
    }

    def __init__(self, file: bool = False):
        super().__init__()
        self.file = file

    # When I wrote this, only God and I understood what I was doing. Now, God only knows
    def format(self, record: logging.LogRecord):
        time = "%(asctime)s" if self.file else colored("%(asctime)s", "magenta")
        level = "%(levelname)s" if self.file else self.FORMATS.get(record.levelno)
        fmt = f"{time} - [{level}][%(name)s]: {record.msg}"
        formatter = logging.Formatter(fmt, datefmt="%Y-%m-%d,%H:%M:%S")
        try:
            return formatter.format(record)

        # Default action of the record.Formatter.format
        except Exception as error:
            record.message = record.getMessage()
            if self.usesTime():
                record.asctime = self.formatTime(record, self.datefmt)
            s = self.formatMessage(record)
            if record.exc_info:
                # Cache the traceback text to avoid converting it multiple times
                # (it's constant anyway)
                if not record.exc_text:
                    record.exc_text = self.formatException(record.exc_info)
            if record.exc_text:
                if s[-1:] != "\n":
                    s = s + "\n"
                s = s + record.exc_text
            if record.stack_info:
                if s[-1:] != "\n":
                    s = s + "\n"
                s = s + self.formatStack(record.stack_info)
            return s


logs_dir = Path(LOGS_DIR)
logs_dir.mkdir(parents=True, exist_ok=True)

logs_file_path = logs_dir / (datetime.utcnow().strftime("%Y_%m_%d") + ".log")
file_handler = logging.FileHandler(logs_file_path, mode="a", encoding="UTF-8")

stream_handler = logging.StreamHandler(stream=sys.stdout)
stream_handler.setFormatter(CustomFormatter())

logging.basicConfig(
    level=logging.DEBUG if DEBUG else logging.INFO,
    handlers=[stream_handler, file_handler],
)

for l in ["aiohttp", "aiohttp.web", "asyncio", "pyrogram"]:
    logging.getLogger(l).setLevel(logging.WARNING)
