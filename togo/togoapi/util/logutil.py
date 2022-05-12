import logging

logging.basicConfig(
    format='%(asctime)s %(levelname)-8s %(message)s',
    level=logging.INFO,
    datefmt='%Y-%m-%d %H:%M:%S')

def errlog(msg):
    logging.error(msg)

def infolog(msg):
    logging.info(msg)