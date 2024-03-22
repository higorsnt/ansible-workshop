import os


class DotEnv:
    @staticmethod
    def smtp_port():
        return os.getenv('SMTP_PORT')

    @staticmethod
    def smtp_server():
        return os.getenv('SMTP_SERVER')

    @staticmethod
    def smtp_login():
        return os.getenv('SMTP_LOGIN')

    @staticmethod
    def smtp_password():
        return os.getenv('SMTP_PASSWORD')

    @staticmethod
    def kafka_server():
        return os.getenv('KAFKA_SERVER')
