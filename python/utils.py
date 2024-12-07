import subprocess
import os


def download_input(url, path):
    if not os.path.exists(path):
        subprocess.run(["curl", "-o", path, url])
