package handler

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sharing-app/internal/service"
	"sharing-app/utils"
	"strconv"
	"strings"
)

type CLIHandler struct {
	userService      *service.UserService
	tripService      *service.TripService
	promotionService *service.PromotionService
}

func NewCLIHandler(userService *service.UserService, tripService *service.TripService, promotionService *service.PromotionService) *CLIHandler {
	return &CLIHandler{
		userService:      userService,
		tripService:      tripService,
		promotionService: promotionService,
	}
}

func (h *CLIHandler) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n==== MAIN MENU ====")
		fmt.Println("1. Users")
		fmt.Println("2. Trips")
		fmt.Println("3. Promotions")
		fmt.Println("4. Analytics")
		fmt.Println("5. Exit")
		fmt.Print("Choose: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			h.usersMenu(reader)
		case "2":
			h.tripsMenu(reader)
		case "3":
			h.promotionsMenu(reader)
		case "4":
			h.analyticsMenu(reader)
		case "5":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (h *CLIHandler) usersMenu(reader *bufio.Reader) {
	for {
		fmt.Println("\n==== USERS MENU ====")
		fmt.Println("1. List users")
		fmt.Println("2. Add user")
		fmt.Println("3. Update user")
		fmt.Println("4. Delete user")
		fmt.Println("5. Back")
		fmt.Print("Choose: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			h.listUsers()
		case "2":
			h.addUser(reader)
		case "3":
			h.updateUser(reader)
		case "4":
			h.deleteUser(reader)
		case "5":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (h *CLIHandler) listUsers() {
	users, err := h.userService.ListUsers(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(users) == 0 {
		fmt.Println("No users found")
		return
	}

	for _, u := range users {
		fmt.Printf(
			"%d | %s | %s | %s\n",
			u.ID,
			u.Fullname,
			u.Phone,
			u.RegistrationDate.Format("2006-01-02"),
		)
	}
}

func (h *CLIHandler) addUser(reader *bufio.Reader) {
	fmt.Print("Fullname: ")
	fullname, _ := reader.ReadString('\n')

	fmt.Print("Phone: ")
	phone, _ := reader.ReadString('\n')

	err := h.userService.CreateUser(
		context.Background(),
		strings.TrimSpace(fullname),
		strings.TrimSpace(phone),
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User created")
}

func (h *CLIHandler) updateUser(reader *bufio.Reader) {
	fmt.Print("User ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	fmt.Print("New fullname: ")
	fullname, _ := reader.ReadString('\n')

	fmt.Print("New phone: ")
	phone, _ := reader.ReadString('\n')

	err = h.userService.UpdateUser(
		context.Background(),
		id,
		strings.TrimSpace(fullname),
		strings.TrimSpace(phone),
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User updated")
}

func (h *CLIHandler) deleteUser(reader *bufio.Reader) {
	fmt.Print("User ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	err = h.userService.DeleteUser(context.Background(), id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User deleted")
}

func (h *CLIHandler) showUserTripStats() {
	stats, err := h.userService.GetTripStats(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(stats) == 0 {
		fmt.Println("No stats found")
		return
	}

	for _, s := range stats {
		fmt.Printf("%s | trips: %d\n", s.UserName, s.TripCount)
	}
}

func (h *CLIHandler) tripsMenu(reader *bufio.Reader) {
	for {
		fmt.Println("\n==== TRIPS MENU ====")
		fmt.Println("1. List trips")
		fmt.Println("2. Create trip")
		fmt.Println("3. Finish trip")
		fmt.Println("4. Cancel trip")
		fmt.Println("5. Delete trip")
		fmt.Println("6. Back")
		fmt.Print("Choose: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			h.listTrips()
		case "2":
			h.createTrip(reader)
		case "3":
			h.finishTrip(reader)
		case "4":
			h.cancelTrip(reader)
		case "5":
			h.deleteTrip(reader)
		case "6":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (h *CLIHandler) listTrips() {
	trips, err := h.tripService.ListTrips(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(trips) == 0 {
		fmt.Println("No trips found")
		return
	}

	for _, t := range trips {
		fmt.Printf(
			"%d | %s | %s | %s | %s\n",
			t.ID,
			t.UserName,
			t.StartParking,
			t.TripStatus,
			t.StartTime.Format("2006-01-02 15:04"),
		)
	}
}

func (h *CLIHandler) createTrip(reader *bufio.Reader) {
	fmt.Print("User ID: ")
	userID := utils.ReadInt(reader)

	fmt.Print("Transport ID: ")
	transportID := utils.ReadInt(reader)

	fmt.Print("Tariff ID: ")
	tariffID := utils.ReadInt(reader)

	fmt.Print("Start Parking ID: ")
	startParkingID := utils.ReadInt(reader)

	err := h.tripService.CreateTrip(
		context.Background(),
		userID,
		transportID,
		tariffID,
		startParkingID,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Trip created")
}

func (h *CLIHandler) finishTrip(reader *bufio.Reader) {
	fmt.Print("Trip ID: ")
	id := utils.ReadInt(reader)

	err := h.tripService.FinishTrip(context.Background(), id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Trip finished")
}

func (h *CLIHandler) cancelTrip(reader *bufio.Reader) {
	fmt.Print("Trip ID: ")
	id := utils.ReadInt(reader)

	fmt.Print("Cancel reason: ")
	reason, _ := reader.ReadString('\n')
	reason = strings.TrimSpace(reason)

	err := h.tripService.CancelTrip(
		context.Background(),
		id,
		reason,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Trip cancelled")
}

func (h *CLIHandler) deleteTrip(reader *bufio.Reader) {
	fmt.Print("Trip ID: ")
	id := utils.ReadInt(reader)

	err := h.tripService.DeleteTrips(context.Background(), id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Trip deleted")
}

func (h *CLIHandler) analyticsMenu(reader *bufio.Reader) {
	for {
		fmt.Println("\n==== ANALYTICS MENU ====")
		fmt.Println("1. Parking stats")
		fmt.Println("2. Trip status stats")
		fmt.Println("3. Trip stats")
		fmt.Println("4. Back")
		fmt.Print("Choose: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			h.showParkingStats()
		case "2":
			h.showTripStatusStats()
		case "3":
			h.showUserTripStats()
		case "4":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (h *CLIHandler) showParkingStats() {
	stats, err := h.tripService.GetParkingStats(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(stats) == 0 {
		fmt.Println("No parking stats found")
		return
	}

	for _, s := range stats {
		fmt.Printf("%s | trips: %d\n", s.Address, s.TripsCount)
	}
}

func (h *CLIHandler) showTripStatusStats() {
	stats, err := h.tripService.GetTripStatusStats(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(stats) == 0 {
		fmt.Println("No trip status stats found")
		return
	}

	for _, s := range stats {
		fmt.Printf("%s | count: %d\n", s.Status, s.Count)
	}
}

func (h *CLIHandler) promotionsMenu(reader *bufio.Reader) {
	for {
		fmt.Println("\n==== PROMOTIONS MENU ====")
		fmt.Println("1. List promotions")
		fmt.Println("2. Create promotion")
		fmt.Println("3. Assign promotion to user")
		fmt.Println("4. Remove promotion from user")
		fmt.Println("5. Show user promotions")
		fmt.Println("6. Back")
		fmt.Print("Choose: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			h.listPromotions()
		case "2":
			h.createPromotion(reader)
		case "3":
			h.assignPromotionToUser(reader)
		case "4":
			h.removePromotionFromUser(reader)
		case "5":
			h.showUserPromotions(reader)
		case "6":
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (h *CLIHandler) listPromotions() {
	promotions, err := h.promotionService.ListPromotions(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(promotions) == 0 {
		fmt.Println("No promotions found")
		return
	}

	for _, p := range promotions {
		fmt.Printf(
			"%d | %s | %d%% | %s - %s | active: %t\n",
			p.ID,
			p.PromoCode,
			p.DiscountPercent,
			p.StartDate.Format("2006-01-02"),
			p.EndDate.Format("2006-01-02"),
			p.IsActive,
		)
	}
}

func (h *CLIHandler) createPromotion(reader *bufio.Reader) {
	fmt.Print("Promo code: ")
	code, _ := reader.ReadString('\n')
	code = strings.TrimSpace(code)

	fmt.Print("Discount percent: ")
	discount := utils.ReadInt(reader)

	err := h.promotionService.CreatePromotion(
		context.Background(),
		code,
		int(discount),
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Promotion created")
}

func (h *CLIHandler) assignPromotionToUser(reader *bufio.Reader) {
	fmt.Print("User ID: ")
	userID := utils.ReadInt(reader)

	fmt.Print("Promotion ID: ")
	promoID := utils.ReadInt(reader)

	err := h.promotionService.AssignToUser(
		context.Background(),
		userID,
		promoID,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Promotion assigned to user")
}

func (h *CLIHandler) removePromotionFromUser(reader *bufio.Reader) {
	fmt.Print("User ID: ")
	userID := utils.ReadInt(reader)

	fmt.Print("Promotion ID: ")
	promoID := utils.ReadInt(reader)

	err := h.promotionService.RemoveFromUser(
		context.Background(),
		userID,
		promoID,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Promotion removed from user")
}

func (h *CLIHandler) showUserPromotions(reader *bufio.Reader) {
	fmt.Print("User ID: ")
	userID := utils.ReadInt(reader)

	promotions, err := h.promotionService.GetUserPromotions(
		context.Background(),
		userID,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(promotions) == 0 {
		fmt.Println("No promotions for this user")
		return
	}

	for _, p := range promotions {
		fmt.Printf(
			"%s | %s | %d%%\n",
			p.UserName,
			p.PromoCode,
			p.Discount,
		)
	}
}
