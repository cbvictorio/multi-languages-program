from django.db import models


# Create your models here.
class Pokemon(models.Model):
    id = models.UUIDField(primary_key=True, null=False, unique=True)
    name = models.TextField(null=False, unique=True)


class User(models.Model):
    id = models.UUIDField(primary_key=True, editable=False, null=False, unique=True)
    favorite_pokemon = models.ForeignKey(Pokemon, on_delete=models.DO_NOTHING)
    team_battle = models.ManyToManyField("BattlePokemon", related_name="pokemons_team")


class BattlePokemon(models.Model):
    user_id = models.ForeignKey(User, on_delete=models.CASCADE, related_name="user")
    pokemon_id = models.ForeignKey(
        Pokemon, on_delete=models.CASCADE, related_name="pokemon"
    )
    alias = models.TextField(null=True, editable=True, max_length=50)
    level: models.IntegerField(default=1, editable=True, null=False)
