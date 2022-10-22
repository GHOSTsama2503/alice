import os
import yaml

locales: list[str] = []
strings: dict[str, str] = {}

def load():
    for locale in os.listdir("locales"):
        language_code = locale.split(".yaml")[0]
        locales.append(language_code)

        with open(os.path.join("locales", locale), "r", encoding="utf-8") as file:
            strings[language_code] = yaml.safe_load(file)
