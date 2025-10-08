from django.db.models import ObjectDoesNotExist
from rest_framework import status
from pokemons.models import Pokemon
from rest_framework.response import Response
from rest_framework.permissions import AllowAny
from pokemons.serializers import PokemonSerializer
from rest_framework.decorators import api_view, permission_classes


@api_view(["GET"])
@permission_classes([AllowAny])
def get_all_pokemons(request):
    pokemons_list_data = Pokemon.objects.all()
    pokemons_list_serializer = PokemonSerializer(pokemons_list_data, many=True)
    response = {"status": 200, "ok": True, "data": pokemons_list_serializer.data}

    return Response(response)


@api_view(["GET"])
@permission_classes([AllowAny])
def get_pokemon_by_id(request, id):
    try:
        pokemon_data = Pokemon.objects.get(id=id)
    except ObjectDoesNotExist:
        return Response(
            {"message": f"pokemon with id: {id} was not found"},
            status=status.HTTP_404_NOT_FOUND,
        )

    serializer = PokemonSerializer(pokemon_data)
    return Response({"ok": True, "data": serializer.data})
