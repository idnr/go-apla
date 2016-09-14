// Copyright 2016 The go-daylight Authors
// This file is part of the go-daylight library.
//
// The go-daylight library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-daylight library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-daylight library. If not, see <http://www.gnu.org/licenses/>.

package tcpserver

import (
	"github.com/DayLightProject/go-daylight/packages/utils"
)

/* Выдаем номер макс. блока
 * запрос шлет демон blocksCollection
 */

func (t *TcpServer) Type10() {
	blockId, err := t.Single("SELECT block_id FROM info_block").Int64()
	if err != nil {
		log.Error("%v", utils.ErrInfo(err))
		return
	}
	_, err = t.Conn.Write(utils.DecToBin(blockId, 4))
	if err != nil {
		log.Error("%v", utils.ErrInfo(err))
		return
	}
}
