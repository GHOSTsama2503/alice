def readable_file_size(size: int, max_decimals: int, unit: str = "B") -> str:
    for prefix in ["", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"]:
        if size < 1024:
            return f"{round(size, max_decimals)} {prefix}{unit}"
        size /= 1024
    return f"{round(size, max_decimals)} Yi{unit}"
    