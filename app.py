import secrets

from flask import Flask, jsonify, render_template, request, redirect, session, flash, get_flashed_messages, abort

from constants import CELL_EMPTY, ORIENTATION_HORIZONTAL, CELL_SHIP
from models import Player, Game

app = Flask(__name__)
app.secret_key = b'Vye=@PW)6LELU5Lg/wW*<%;U&}NV#h*!U{H{ZPkwQLQj'

players = {}
games = {}


@app.route('/', methods=['GET'])
def index():
    if 'token' not in session or session['token'] not in players:
        return redirect('/signup')

    token = session['token']
    join_game(players[token])

    return render_template('index.html')


@app.route('/signup', methods=['GET', 'POST'])
def signup():
    if request.method == 'POST':
        player_name = request.form['player_name']

        if player_name == '':
            flash('empty_player_name')
            return redirect('/signup')

        while True:
            token = secrets.token_hex(64)
            if token not in players:
                break
        players[token] = Player(name=player_name, active_game=None)
        session['token'] = token

        return redirect('/')

    return render_template('register.html', errors=get_flashed_messages())


@app.route('/status')
def status():
    if 'token' not in session or session['token'] not in players:
        return abort(401)

    token = session['token']
    player = players[token]

    return jsonify({
        'game': games[player.active_game].to_json(),
        'player': player.to_json(),
        'battlefield': games[player.active_game].get_battlefield(player)
    })


@app.route('/action', methods=['POST'])
def ready():
    if 'token' not in session or session['token'] not in players:
        return abort(401)

    token = session['token']
    player = players[token]

    act = request.args.get('act')
    if act == 'ready':
        games[player.active_game].player_ready(player, request.get_json())
    elif act == 'shot':
        if not games[player.active_game].process_shot(player, request.form.get('target')):
            return abort(403)
    elif act == 'surround':
        if not games[player.active_game].surround(player):
            return abort(403)
    else:
        return abort(400)

    return jsonify({
        'game': games[player.active_game].to_json(),
        'player': player.to_json(),
        'battlefield': games[player.active_game].get_battlefield(player)
    })


@app.route('/battlefield', methods=['POST'])
def generate_field():
    field = [[]] * 10
    for i in range(10):
        field[i] = [CELL_EMPTY] * 10

    ships = [1, 2, 3, 4]
    sizes = [4, 3, 2, 1]
    current_ship = 0

    while True:
        while True:
            orientation = 0 if secrets.randbelow(100) < 50 else 1
            ship_x, ship_y = secrets.randbelow(10), secrets.randbelow(10)

            if ship_could_be_placed_here(field, sizes[current_ship], orientation, ship_x, ship_y):
                break

        if orientation == ORIENTATION_HORIZONTAL:
            for x in range(ship_x, ship_x + sizes[current_ship]):
                field[ship_y][x] = CELL_SHIP
        else:
            for y in range(ship_y, ship_y + sizes[current_ship]):
                field[y][ship_x] = CELL_SHIP

        ships[current_ship] -= 1
        if ships[current_ship] == 0:
            current_ship += 1

        if current_ship >= len(ships):
            break

    return jsonify(field)


def ship_could_be_placed_here(field, size, orientation, x, y):
    if x + size - 1 > 9:
        return False

    if y + size - 1 > 9:
        return False

    if orientation == ORIENTATION_HORIZONTAL:
        for i in range(-1 if x > 0 else 0, size + 1):
            if x + i > 9:
                continue

            if y > 0 and field[y - 1][x + i] != CELL_EMPTY:
                return False

            if field[y][x + i] != CELL_EMPTY:
                return False

            if y < 9 and field[y + 1][x + i] != CELL_EMPTY:
                return False
    else:
        for i in range(-1 if y > 0 else 0, size + 1):
            if y + i > 9:
                continue

            if x > 0 and field[y + i][x - 1] != CELL_EMPTY:
                return False

            if field[y + i][x] != CELL_EMPTY:
                return False

            if x < 9 and field[y + i][x + 1] != CELL_EMPTY:
                return False

    return True


def join_game(player):
    if player.active_game is not None and games[player.active_game].winner is None:
        return

    game = find_game()
    if game is None:
        game = Game(player1=player, player2=None)
        games[game.uuid] = game
    else:
        game.join(player)

    player.active_game = game.uuid


def find_game():
    for game in games.values():
        if game.state == Game.STATE_NO_OPPONENT:
            return game

    return None


if __name__ == '__main__':
    app.run()
