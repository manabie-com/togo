import os
import sys

if os.path.exists("/togo"):
    sys.path.insert(0, os.path.abspath("/togo"))
    print("Imported the modules!")

togo_abspath = os.path.join(os.path.dirname(__file__), "../../../")
sys.path.insert(0, os.path.abspath(togo_abspath))

    