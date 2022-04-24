/**
 * @Author: chentong
 * @Date: 2022/04/24 23:29
 */

package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image

func loadIcon(name string) image.Image {
	return nil
}

// Icon 并发不安全
func Icon(name string) image.Image {
	icon, ok := icons[name]
	if !ok {
		icon := loadIcon(name)
		icons[name] = icon
	}
	return icon
}

/*  *   * * * * * * * * * * * * * * * * * * *   *         */

var loadIconsOnce sync.Once

func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// IconOnec 并发安全
func IconOnec(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]

}
