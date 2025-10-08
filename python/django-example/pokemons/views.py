from rest_framework.decorators import api_view, permission_classes
from rest_framework.response import Response
from pokemons.models import Pokemon
from pokemons.serializers import PokemonSerializer
from rest_framework.permissions import AllowAny


@api_view(["GET"])
@permission_classes([AllowAny])
def get_all_pokemons(request):
    pokemons_list_data = Pokemon.objects.all()
    charizard_data = Pokemon.objects.get(id=2)

    pokemons_list_serializer = PokemonSerializer(pokemons_list_data, many=True)
    charizard_data_serializer = PokemonSerializer(charizard_data)

    response = {
        "status": 200,
        "ok": True,
        "error": False,
        "list": pokemons_list_serializer.data,
        "entity": charizard_data_serializer.data,
    }

    return Response(response)


@api_view(["POST"])
@permission_classes([AllowAny])
def create_initial_pokemon_data(request):
    pokemons_list = []
    return Response(pokemons_list)
