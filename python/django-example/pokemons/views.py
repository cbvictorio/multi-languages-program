from django.shortcuts import render
from pokemons.models import Pokemon

# Create your views here.


def pokemons_view(request):
    pokemons_list: list[Pokemon] = [{"id": 1, "name": "fakin charizard"}]
    context = {"pokemons_list": pokemons_list}
    return render(request, "pokemons/pokemons_list.html", context)
