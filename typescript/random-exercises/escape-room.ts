export class EscapeRoom {
  numberOfPlayers: number;
  numberOfRooms: number;
  gameStatus: Record<number, number> = {};

  constructor(numberOfPlayers: number, numberOfRooms) {
    this.numberOfPlayers = numberOfPlayers;
    this.numberOfRooms = numberOfRooms;

    const playersNumberArray: number[] = Array.from(
      new Array(numberOfPlayers)
    ).map((_, i) => i + 1);

    for (const number of playersNumberArray) {
      this.gameStatus[number] = 1;
    }
  }

  movePlayerToNextStage(playerNumber: number) {
    const currentPlayerRoom = this.gameStatus[playerNumber];

    if (currentPlayerRoom === this.numberOfRooms) {
      return console.log(`Player no.${playerNumber} has won!`);
    }

    this.gameStatus[playerNumber] = currentPlayerRoom + 1;
  }

  printRanking() {
    const ranking = Object.entries(this.gameStatus).sort((a, b) => b[1] - a[1]);
    let i = 1;

    for (const [player, room] of ranking) {
      console.log(`Place (${i}) => Player ${player}, Room ${room}`);
      i++;
    }
  }

  print() {
    console.log(this.gameStatus);
  }
}

const escapeRoom = new EscapeRoom(4, 4);
escapeRoom.movePlayerToNextStage(1);
escapeRoom.movePlayerToNextStage(1);
escapeRoom.movePlayerToNextStage(1);
// escapeRoom.print();
escapeRoom.printRanking();
