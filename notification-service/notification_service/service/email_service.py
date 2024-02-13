import ssl
from email.mime.image import MIMEImage
from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

import chevron
import smtplib

from notification_service.dot_env import DotEnv
from notification_service.model.order import Order
from notification_service.model.user import User


def send_email(template, user: User):
    sender_email = DotEnv.smtp_login()
    receiver_email = user.email
    message = MIMEMultipart("alternative")
    html = MIMEText(template, "html")

    message["Subject"] = "Order Confirmation"
    message["From"] = sender_email
    message["To"] = receiver_email
    message.attach(html)

    with open('../template/resources/logo.png', 'rb') as logo:
        logo_img = MIMEImage(logo.read())
        logo_img.add_header('Content-ID', '<logo>')
        message.attach(logo_img)

    with open('../template/resources/success.png', 'rb') as success:
        success_img = MIMEImage(success.read())
        success_img.add_header('Content-ID', '<success>')
        message.attach(success_img)

    context = ssl.create_default_context()

    with smtplib.SMTP_SSL(DotEnv.smtp_server(), DotEnv.smtp_port(), context=context) as server:
        server.login(DotEnv.smtp_login(), DotEnv.smtp_password())
        server.sendmail(sender_email, receiver_email, message.as_string())


def render_template(template_path: str, args: dict):
    with template_path as f:
        return chevron.render(template=f, data=args)


def send_order_confirmation_email(order: Order):
    total_price = sum((p.quantity * p.price) for p in order.products)
    billing_address = f'{order.company.address.street}, {order.company.address.number}, {order.company.address.city} - {order.company.address.state}'
    shipping_address = f'{order.user.address.street}, {order.user.address.number}, {order.user.address.city} - {order.user.address.state}'

    args = {
        'order': {
            'id': order.id,
            'items': [{'name': p.name, 'price': p.price, 'quantity': p.quantity} for p in order.products],
            'total_price': total_price,
            'unsubscribe': 'http://example.com/unsubscribe'
        },
        'company': {
            'name': order.company.name,
            'billing_address': billing_address
        },
        'user': {
            'shipping_address': shipping_address
        },
    }
    email_body = render_template('../template/order_confirmation.mustache', args)
    send_email(email_body, order.user)
