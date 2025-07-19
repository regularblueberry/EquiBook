package application




type AppError struct {
	System AppSystemID;
	Code error;
	Message string;
}


func (self *App) TryRecover(err AppError){

	switch err.System {
	case SYS_RENDER:
	case SYS_NETWORK:
	case SYS_STORAGE:
	default:
	}


}
