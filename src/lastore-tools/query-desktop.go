/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"fmt"

	"internal/querydesktop"

	"github.com/codegangsta/cli"
)

var CMDQueryDesktop = cli.Command{
	Name:  "querydesktop",
	Usage: `pkgname`,
	Action: func(ctx *cli.Context) error {
		querydesktop.InitDB()
		if ctx.NArg() != 1 {
			_ = cli.ShowAppHelp(ctx)
			return nil
		}
		fmt.Println(querydesktop.QueryDesktopFile(ctx.Args()[0]))
		return nil
	},
}
