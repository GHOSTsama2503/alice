import logging
from datetime import datetime, timezone
from termcolor import colored
import os

class UserFormatter(logging.Formatter):
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

    def format(self, record):
        time = "%(asctime)s" if self.file else colored("%(asctime)s", "magenta")
        level = "%(levelname)s" if self.file else self.FORMATS.get(record.levelno)
        fmt = f"{time} - [{level}] {record.msg}"
        formatter = logging.Formatter(fmt, datefmt = "%H:%M:%S")
        return formatter.format(record)

log = logging.getLogger("logger")
logger_level = logging.DEBUG

def log_file_path() -> str:
    dt = datetime.now(timezone.utc)
    pfmt = dt.strftime("./logs/%Y/%B")
    os.makedirs(pfmt, exist_ok = True)
    return pfmt + dt.strftime("/%d.log")

file_handler = logging.FileHandler(log_file_path())
file_handler.setLevel(logger_level)
file_handler.setFormatter(UserFormatter(True))

stream_handler = logging.StreamHandler()
stream_handler.setLevel(logger_level)
stream_handler.setFormatter(UserFormatter())

log.addHandler(file_handler)
log.addHandler(stream_handler)
log.setLevel(logger_level)