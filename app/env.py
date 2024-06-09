from datetime import datetime
from os import getenv
from pathlib import Path

from dotenv import load_dotenv

load_dotenv()


def as_bool(value: str) -> bool:
    return value.lower() in ("1", "true", "t", "yes", "y")


DEBUG = as_bool(getenv("DEBUG", "0"))
START_TIME = datetime.utcnow().timestamp()

# Telegram Client:
API_ID = getenv("API_ID")
API_HASH = getenv("API_HASH")
BOT_TOKEN = getenv("BOT_TOKEN")
BOT_SESSION = getenv("BOT_SESSION")

# Telegram Client Settings:
BOT_PICTURE = getenv("BOT_PICTURE", "https://telegra.ph/file/fdc88663a5eb402f46d84.jpg")
CLIENT_NAME = getenv("CLIENT_NAME", "alice")
START_MESSAGE = as_bool(getenv("START_MESSAGE", "1"))

# Telegram:
ADMIN_ID = int(getenv("ADMIN_ID"))

# Custom Settings:
DEFAULT_LANG = getenv("DEFAULT_LANG", "es")
PREFIXES = list(getenv("PREFIXES", "/"))

# Paths
LOCALES_DIR = getenv("LOCALES_DIR", "locales/")
PLUGINS_PATH = getenv("PLUGINS_PATH", "app/handlers/")
SESSIONS_PATH = getenv("SESSIONS_PATH", "sessions/")
STORAGE_PATH = getenv("STORAGE_PATH", "storage/")
TEMP_PATH = Path(getenv("TEMP_PATH", "/tmp/alice"))

THUMBNAIL = Path(getenv("THUMBNAIL", "thumb.jpg"))
THUMBNAIL = THUMBNAIL if THUMBNAIL.exists() else None
