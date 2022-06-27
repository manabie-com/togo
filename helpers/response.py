from abc import ABC, abstractmethod


class AbstractResponse(ABC):
    @abstractmethod
    def keys(self):
        raise NotImplementedError

    @abstractmethod
    def __getitem__(self, key):
        raise NotImplementedError
