# Generated by Django 3.1.3 on 2020-11-26 15:38

from django.db import migrations
import phonenumber_field.modelfields


class Migration(migrations.Migration):

    dependencies = [
        ('user', '0002_auto_20201126_1632'),
    ]

    operations = [
        migrations.AlterField(
            model_name='profile',
            name='phone',
            field=phonenumber_field.modelfields.PhoneNumberField(default='000000000000', max_length=128, region=None),
        ),
    ]
