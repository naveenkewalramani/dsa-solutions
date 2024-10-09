class ParkingLot():
    def __init__(self, floors, carSpaces, truckSpaces, bikeSpaces):
        self.parkingLot = {}
        self.floorCount = floors
        cars = [False for j in range(0, carSpaces)]
        trucks =  [False for j in range(0, truckSpaces)]
        bikes =  [False for j in range(0, bikeSpaces)]
        for i in range(0, floors):
            self.parkingLot[i] = [cars, trucks, bikes]
    
    def getFreeSpaces(self):
        for floor, spots in self.parkingLot.items():
            cars = []
            for j in range(0, len(spots[0])):
                if spots[0][j] == False:
                    cars.append(j)
            print("Floor: ", floor, "Free Spaces cars: ", cars)
            trucks = []
            for j in range(0, len(spots[1])):
                if spots[1][j] == False:
                    trucks.append(j)
            print("Floor: ", floor, "Free Spaces trucks: ",trucks)
            bikes = []
            for j in range(0, len(spots[2])):
                if spots[2][j] == False:
                    bikes.append(j)
            print("Floor: ", floor, "Free Spaces bikes: ",bikes)
        
    def alotSpace(self, type):
        if type == 0:
            for floor, spots in self.parkingLot.items():
                for j in range(0, len(spots[0])):
                    if spots[0][j] == False:
                        spots[0][j] = True
                        self.parkingLot[floor] = spots
                        print("Space alloted: floor: ", floor, " and spot: ", j)
                        return True
            print("No spot of cars")
            return False
        if type == 1:
            for floor, spots in self.parkingLot.items():
                for j in range(0, len(spots[1])):
                    if spots[1][j] == False:
                        spots[1][j] = True
                        self.parkingLot[floor] = spots
                        print("Space alloted: floor: ", floor, " and spot: ", j)
                        return True
            print("No spot of trucks")
            return False
        if type == 2:
            for floor, spots in self.parkingLot.items():
                for j in range(0, len(spots[2])):
                    if spots[2][j] == False:
                        spots[2][j] = True
                        self.parkingLot[floor] = spots
                        print("Space alloted: floor: ", floor, " and spot: ", j)
                        return True
            print("No spot of bikes")
            return False


obj = ParkingLot(3, 5, 5, 4)
obj.getFreeSpaces()
obj.alotSpace(0)
obj.alotSpace(1)
obj.alotSpace(2)
obj.alotSpace(2)
obj.alotSpace(1)
obj.getFreeSpaces()