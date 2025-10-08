from django.urls import path
from . import views

# application urls

urlpatterns = [
    path("", views.get_all_pokemons),
    path("<int:id>/", views.get_pokemon_by_id),
]
