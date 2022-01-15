#!/usr/bin/env python3
import os
import socket
import binascii
import unittest
import time
import requests
import hashlib
from faker import Faker

HOST = "localhost"
if os.getenv("HOST") != None:
  HOST = os.getenv("HOST")

class TestToDo(unittest.TestCase):
  @classmethod
  def setUpClass(cls):
    cls.server = f"http://{HOST}:3000"
    cls.faker = Faker()

  def get_fresh_key(self):
    user_id = binascii.hexlify(os.urandom(10)).decode('ascii')
    date = self.faker.date_time_between(start_date='-2y', end_date='now').strftime("%Y%m%d")
    task_key = binascii.hexlify(os.urandom(10)).decode('ascii')
    return f"{user_id}/{date}/{task_key}"

  def get_key_with_host(self):
    return f"{self.server}/{self.get_fresh_key()}"

  def test_putgetdelete(self):
    remote = self.get_key_with_host()
 
    r = requests.put(remote, data="onyou")
    self.assertEqual(r.status_code, 201)

    r = requests.get(remote)
    self.assertEqual(r.status_code, 200)
    self.assertEqual(r.text, "onyou")

    r = requests.delete(remote)
    self.assertEqual(r.status_code, 204)

    r = requests.get(remote)
    self.assertEqual(r.status_code, 404)

  def test_doubledelete(self):
    remote = self.get_key_with_host()

    r = requests.put(remote, data="onyou")
    self.assertEqual(r.status_code, 201)

    r = requests.delete(remote)
    self.assertEqual(r.status_code, 204)

    r = requests.delete(remote)
    self.assertNotEqual(r.status_code, 204)

  def test_doubleput(self):
    remote = self.get_key_with_host()
    r = requests.put(remote, data="onyou")
    self.assertEqual(r.status_code, 201)

    r = requests.put(remote, data="onyou2")
    self.assertEqual(r.status_code, 204)

    r = requests.get(remote)
    self.assertEqual(r.status_code, 200)
    self.assertEqual(r.text, "onyou2")

  def test_doubleputdelete(self):
    remote = self.get_key_with_host()

    r = requests.put(remote, data="onyou")
    self.assertEqual(r.status_code, 201)

    r = requests.delete(remote)
    self.assertEqual(r.status_code, 204)

    r = requests.put(remote, data="onyou")
    self.assertEqual(r.status_code, 201)

  def test_10keys(self):
    keys = [self.get_key_with_host() for i in range(10)]

    for k in keys:
      r = requests.put(k, data=hashlib.md5(k.encode('ascii')).hexdigest())
      self.assertEqual(r.status_code, 201)

    for k in keys:
      r = requests.get(k)
      self.assertEqual(r.status_code, 200)
      self.assertEqual(r.text, hashlib.md5(k.encode('ascii')).hexdigest())

    for k in keys:
      r = requests.delete(k)
      self.assertEqual(r.status_code, 204)

  def test_nonexistent_key(self):
    remote = self.get_key_with_host()
    r = requests.get(remote)
    self.assertEqual(r.status_code, 404)

  def test_large_key(self):
    remote = self.get_key_with_host()
    data = b"a"*(16*1024*1024)

    r = requests.put(remote, data=data)
    self.assertEqual(r.status_code, 201)

    r = requests.get(remote)
    self.assertEqual(r.status_code, 200)
    self.assertEqual(r.content, data)

    r = requests.delete(remote)
    self.assertEqual(r.status_code, 204)

  def test_limit(self):
    #this key 123/20220114 have limit 1 task
    remote = self.server + "/123/20220114/"

    r = requests.put(remote + "task1", data="task1")
    self.assertEqual(r.status_code, 201)

    r = requests.get(remote + "task1")
    self.assertEqual(r.status_code, 200)
    self.assertEqual(r.content, b"task1")

    r = requests.put(remote + "task2", data="task2")
    self.assertEqual(r.status_code, 403)
    r = requests.get(remote + "task2")
    self.assertNotEqual(r.status_code, 200)

    r = requests.delete(remote + "task1")
    self.assertEqual(r.status_code, 204)

if __name__ == '__main__':
  # wait for servers
  while 1:
    try:
      s = socket.create_connection((f"{HOST}", 3000), timeout=0.5)
      s.close()
      break
    except (ConnectionRefusedError, OSError):
      time.sleep(0.5)
      print("waiting for servers")
      continue

  unittest.main()
