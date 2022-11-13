package service

import "tool-server/internal/global"

func GetChangelog() global.Release {
	return global.ChangeLog
}

func GetNavigation() global.Navigation {
	return global.NavigationMenu
}
