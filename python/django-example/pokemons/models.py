from django.db import models


# Create your models here.
class Pokemon(models.Model):
    id = models.UUIDField(primary_key=True, editable=False, null=False, unique=True)
    name = models.TextField(null=False)
