import secrets
from datetime import datetime
from uuid import uuid4

from constants import CELL_SHIP


class Player(object):
    def __init__(self, name, active_game):
        self.uuid = uuid4()
        self.name = name
        self.active_game = active_game

    def to_json(self):
        return {'uuid': self.uuid, 'name': self.name, 'active_game': self.active_game}


class Shot(object):
    def __init__(self, target, hit_on_target, hidden):
        self.target = target
        self.hit_on_target = hit_on_target
        self.hidden = hidden
        self.time = datetime.utcnow()

    def to_json(self):
        return {'target': self.target, 'hit_on_target': self.hit_on_target, 'time': self.time}


class Game(object):
    STATE_NO_OPPONENT = -1
    STATE_OPPONENT_FOUND = 4

    STATE_PLAYER1 = 0
    STATE_PLAYER1_READY = 1
    STATE_PLAYER2 = 2
    STATE_PLAYER2_READY = 3

    def __init__(self, player1, player2):
        self.uuid = uuid4()
        self.state = Game.STATE_NO_OPPONENT
        self.player1 = player1
        self.player1_field = []
        self.player1_shots = []
        self.player1_hits = []
        self.player1_ships = []
        self.player2 = player2
        self.player2_field = []
        self.player2_shots = []
        self.player2_hits = []
        self.player2_ships = []
        self.winner = None
        self.log = []
        self.last_action = None

    def join(self, player):
        self.player2 = player
        self.state = Game.STATE_OPPONENT_FOUND
        self.last_action = datetime.utcnow()

    def get_battlefield(self, player):
        return self.player1_field if self._player_is_the_first_one(player) else self.player2_field

    def player_ready(self, player, field):
        if self._player_is_the_first_one(player):
            if self.state == Game.STATE_OPPONENT_FOUND:
                self.state = Game.STATE_PLAYER1_READY
            elif self.state == Game.STATE_PLAYER2_READY:
                self.state = secrets.choice([Game.STATE_PLAYER1, Game.STATE_PLAYER2])
            else:
                return

            self.player1_field = field
            self.player1_ships = Game._find_ships(field)
        else:
            if self.state == Game.STATE_OPPONENT_FOUND:
                self.state = Game.STATE_PLAYER2_READY
            elif self.state == Game.STATE_PLAYER1_READY:
                self.state = secrets.choice([Game.STATE_PLAYER1, Game.STATE_PLAYER2])
            else:
                return

            self.player2_field = field
            self.player2_ships = Game._find_ships(field)

    def process_shot(self, player, target):
        if self._player_is_the_first_one(player):
            if self.state != Game.STATE_PLAYER1:
                return False
        else:
            if self.state != Game.STATE_PLAYER2:
                return False

        field = self.player2_field if self._player_is_the_first_one(player) else self.player1_field
        hits = self.player1_hits if self._player_is_the_first_one(player) else self.player2_hits
        ships = self.player2_ships if self._player_is_the_first_one(player) else self.player1_ships
        player_shots = self.player1_shots if self._player_is_the_first_one(player) else self.player2_shots

        if Game._already_processed(player_shots, target):
            return False

        x, y = target.split(':')
        x, y = int(x), int(y)

        hit_on_target = field[y][x] == CELL_SHIP
        shot = Shot(target, hit_on_target, False)
        player_shots.append(shot)

        if hit_on_target:
            hits.append(target)

        if not hit_on_target:
            self.state = Game.STATE_PLAYER2 if self._player_is_the_first_one(player) else Game.STATE_PLAYER1
        elif Game._ship_fully_destroyed(ships, x, y, hits):
            ship = Game._get_attacked_ship(ships, x, y)
            for point in ship:
                ship_x, ship_y = [int(c) for c in point.split(':')]

                if (ship_x > 0 and field[ship_y][ship_x - 1] != CELL_SHIP
                        and not Game._shot_already_recorded(player_shots, ship_x - 1, ship_y)):
                    player_shots.append(Shot('{}:{}'.format(ship_x - 1, ship_y), False, True))
                if (ship_x < 9 and field[ship_y][ship_x + 1] != CELL_SHIP
                        and not Game._shot_already_recorded(player_shots, ship_x + 1, ship_y)):
                    player_shots.append(Shot('{}:{}'.format(ship_x + 1, ship_y), False, True))
                if (ship_y > 0 and field[ship_y - 1][ship_x] != CELL_SHIP
                        and not Game._shot_already_recorded(player_shots, ship_x, ship_y - 1)):
                    player_shots.append(Shot('{}:{}'.format(ship_x, ship_y - 1), False, True))
                if (ship_y < 9 and field[ship_y + 1][ship_x] != CELL_SHIP
                        and not Game._shot_already_recorded(player_shots, ship_x, ship_y + 1)):
                    player_shots.append(Shot('{}:{}'.format(ship_x, ship_y + 1), False, True))

                if (ship_x > 0 and ship_y > 0 and field[ship_y - 1][ship_x - 1] != CELL_SHIP
                        and not Game._shot_already_recorded(player_shots, ship_x - 1, ship_y - 1)):
                    player_shots.append(Shot('{}:{}'.format(ship_x - 1, ship_y - 1), False, True))
                if (ship_x < 9 and ship_y < 9 and field[ship_y + 1][ship_x + 1] != CELL_SHIP
                        and not Game._shot_already_recorded(player_shots, ship_x + 1, ship_y + 1)):
                    player_shots.append(Shot('{}:{}'.format(ship_x + 1, ship_y + 1), False, True))
                if (ship_x < 9 and ship_y > 0 and field[ship_y - 1][ship_x + 1] != CELL_SHIP
                        and not Game._shot_already_recorded(player_shots, ship_x + 1, ship_y - 1)):
                    player_shots.append(Shot('{}:{}'.format(ship_x + 1, ship_y - 1), False, True))
                if (ship_x > 0 and ship_y < 9 and field[ship_y + 1][ship_x - 1] != CELL_SHIP
                        and not Game._shot_already_recorded(player_shots, ship_x - 1, ship_y + 1)):
                    player_shots.append(Shot('{}:{}'.format(ship_x - 1, ship_y + 1), False, True))

            print('Fully destroyed', ship)

        if len(hits) == 20:
            self.winner = player.uuid

        return True

    def surround(self, player):
        if self.winner is not None:
            return False

        self.winner = self.player2.uuid if self._player_is_the_first_one(player) else self.player1.uuid
        return True

    def _player_is_the_first_one(self, player):
        return player.uuid == self.player1.uuid

    @staticmethod
    def _already_processed(shots, target):
        return Game._shot_already_recorded(shots, *[int(c) for c in target.split(':')])

    @staticmethod
    def _ship_fully_destroyed(ships, x, y, hits):
        ship = Game._get_attacked_ship(ships, x, y)
        if ship is None:
            return False
        return all([hit in hits for hit in ship])

    @staticmethod
    def _get_attacked_ship(ships, x, y):
        for ship in ships:
            if '{}:{}'.format(x, y) in ship:
                return ship

        return None

    @staticmethod
    def _shot_already_recorded(shots, x, y):
        for shot in shots:
            if shot.target == '{}:{}'.format(x, y):
                return True
        return False

    @staticmethod
    def _find_ships(field):
        ships = []
        ship = []

        processed = []
        for y in range(10):
            for x in range(10):
                if '{}:{}'.format(x, y) in processed:
                    continue

                if field[y][x] == CELL_SHIP:
                    ship.append('{}:{}'.format(x, y))
                    processed.append('{}:{}'.format(x, y))

                    if x < 9 and field[y][x + 1] == CELL_SHIP:
                        ship.append('{}:{}'.format(x + 1, y))
                        processed.append('{}:{}'.format(x + 1, y))

                    if x < 8 and field[y][x + 1] == CELL_SHIP and field[y][x + 2] == CELL_SHIP:
                        ship.append('{}:{}'.format(x + 2, y))
                        processed.append('{}:{}'.format(x + 2, y))

                    if (x < 7 and field[y][x + 1] == CELL_SHIP and field[y][x + 2] == CELL_SHIP
                            and field[y][x + 3] == CELL_SHIP):
                        ship.append('{}:{}'.format(x + 3, y))
                        processed.append('{}:{}'.format(x + 3, y))

                    if y < 9 and field[y + 1][x] == CELL_SHIP:
                        ship.append('{}:{}'.format(x, y + 1))
                        processed.append('{}:{}'.format(x, y + 1))

                    if y < 8 and field[y + 1][x] == CELL_SHIP and field[y + 2][x] == CELL_SHIP:
                        ship.append('{}:{}'.format(x, y + 2))
                        processed.append('{}:{}'.format(x, y + 2))

                    if (y < 7 and field[y + 1][x] == CELL_SHIP and field[y + 2][x] == CELL_SHIP
                            and field[y + 3][x] == CELL_SHIP):
                        ship.append('{}:{}'.format(x, y + 3))
                        processed.append('{}:{}'.format(x, y + 3))
                elif len(ship):
                    ships.append(ship)
                    ship = []

        return ships

    def to_json(self):
        return {
            'player1': {
                'uuid': self.player1.uuid,
                'name': self.player1.name,
                'shots': [x.to_json() for x in self.player1_shots],
            },
            'player2': {
                'id': self.player2.uuid if self.player2 is not None else None,
                'name': self.player2.name if self.player2 is not None else None,
                'shots': [x.to_json() for x in self.player2_shots],
            },
            'log': self.log,
            'state': self.state,
            'winner': self.winner,
        }
