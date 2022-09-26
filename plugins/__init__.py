# Import all the py/pyc files in this folder dinamically
# Do not modify unless you know what you're doing

from importlib import import_module
from pathlib import Path


path = Path(__file__).parent.relative_to(Path(".").absolute())
for file in path.iterdir():
    if not file.is_file():
        continue
    if file.name.lower().startswith("__"):
        continue
    if not file.name.lower().endswith((".py", ".pyc")):
        continue
    if file.name.count(".") > 1:
        continue

    module = "." + file.name.replace(".pyc", "").replace(".py", "")

    import_module(module, __package__)

del path
del file
del module
del Path
del import_module
