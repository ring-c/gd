// Package Rect2i provides a 2D axis-aligned bounding box using integer coordinates.
package Rect2i

import (
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Int"
	"graphics.gd/variant/Vector2i"
)

type Side int

const (
	SideLeft   Side = 0
	SideTop    Side = 1
	SideRight  Side = 2
	SideBottom Side = 3
)

// PositionSize represents an axis-aligned rectangle in a 2D space, using integer coordinates.
// It is defined by its position and size, which are [Vector2i.XY]. Because it does not rotate,
// it is frequently used for fast overlap tests (see intersects).
//
// For floating-point coordinates, see Rect2.
//
// Note: Negative values for size are not supported. With negative size, most functions do
// not work correctly. Use [Abs] to get an equivalent Rect2i with a non-negative size.
type PositionSize = struct {
	Position Vector2i.XY // The origin point. This is usually the top-left corner of the rectangle.

	// The rectangle's width and height, starting from position. Setting this value also affects the
	// end point.
	//
	// Note: It's recommended setting the width and height to non-negative values, as most functions
	// assume that the position is the top-left corner, and the end is the bottom-right corner. To
	// get an equivalent rectangle with non-negative size, use [Abs].
	Size Vector2i.XY
}

// New constructs a [PositionSize] by setting its position to (x, y), and its size to (w, h).
func New[X Float.Any | Int.Any](x, y, w, h X) PositionSize { //gd:Rect2i(x:float,y:float,w:float,h:float)
	return PositionSize{Vector2i.New(x, y), Vector2i.New(w, h)}
}

// End returns the ending point. This is usually the bottom-right corner of the rectangle, and is
// equivalent to position + size.
func End(rect PositionSize) Vector2i.XY {
	return Vector2i.Add(rect.Position, rect.Size)
}

// Abs returns a [PositionSize] equivalent to this rectangle, with its width and height
// modified to be non-negative values, and with its position being the top-left
// corner of the rectangle.
//
// Note: It's recommended to use this method when size is negative, as most functions
// will assume that the position is the top-left corner, and the end is the bottom-right corner.
func Abs(v PositionSize) PositionSize { //gd:Rect2i.abs
	return PositionSize{Vector2i.Add(v.Position, Vector2i.Mini(v.Size, 0)), Vector2i.Abs(v.Size)}
}

// Inside returns true if the rectangle is enclosed by the enclosure rectangle.
func Inside(enclosure, rect PositionSize) bool { //gd:Rect2i.encloses
	a := enclosure
	b := rect
	return (b.Position.X >= a.Position.Y) && (b.Position.Y >= a.Position.Y) &&
		((b.Position.X + b.Size.X) <= (a.Position.X + a.Size.X)) &&
		((b.Position.Y + b.Size.Y) <= (a.Position.Y + a.Size.Y))
}

// ExpandTo returns a copy of this rectangle expanded to align the edges with the
// given to point, if necessary.
func ExpandTo(point Vector2i.XY, rect PositionSize) PositionSize { //gd:Rect2i.expand
	var begin = rect.Position
	var end = Vector2i.Add(rect.Position, rect.Size)
	if point.X < begin.X {
		begin.X = point.X
	}
	if point.Y < begin.Y {
		begin.Y = point.Y
	}
	if point.X > end.X {
		end.X = point.X
	}
	if point.Y > end.Y {
		end.Y = point.Y
	}
	return PositionSize{
		Position: begin,
		Size:     Vector2i.Sub(end, begin),
	}
}

// Area returns the rectangle's area. This is equivalent to
//
//	size.X * size.Y
//
// See also [HasArea].
func Area(rect PositionSize) int { //gd:Rect2i.get_area
	return int(rect.Size.X) * int(rect.Size.Y)
}

// Center returns the center point of the rectangle position + (size / 2.0).
func Center(rect PositionSize) Vector2i.XY { //gd:Rect2i.get_center
	return Vector2i.Add(rect.Position, Vector2i.DivX(rect.Size, 2))
}

// Expand returns a copy of this rectangle extended on all sides by the given
// amount. A negative amount shrinks the rectangle instead. See also
// [ExpandSides] and [ExpandSide].
func Expand[X Int.Any](rect PositionSize, amount X) PositionSize { //gd:Rect2i.grow
	return PositionSize{
		Position: Vector2i.SubX(rect.Position, amount),
		Size:     Vector2i.AddX(rect.Size, amount*2),
	}
}

// ExpandSides returns a copy of this rectangle with its left, top, right, and bottom
// sides extended by the given amounts. Negative values shrink the sides, instead.
// See also [Expand] and [ExpandSide].
func ExpandSides[X Int.Any](rect PositionSize, left, top, right, bottom X) PositionSize { //gd:Rect2i.grow_individual
	rect.Position.X -= int32(left)
	rect.Position.Y -= int32(top)
	rect.Size.X += int32(left + right)
	rect.Size.Y += int32(top + bottom)
	return rect
}

// ExpandSize returns a copy of this rectangle with its side extended by the given amount
// (see Side constants). A negative amount shrinks the rectangle, instead. See also
// [Expand] and [ExpandSides].
func ExpandSide[X Float.Any | Int.Any](rect PositionSize, side Side, amount X) PositionSize { //gd:Rect2i.grow_side
	switch side {
	case SideLeft:
		rect.Position.X -= int32(amount)
		rect.Size.X += int32(amount)
	case SideTop:
		rect.Position.Y -= int32(amount)
		rect.Size.Y += int32(amount)
	case SideRight:
		rect.Size.X += int32(amount)
	case SideBottom:
		rect.Size.Y += int32(amount)
	}
	return rect
}

// HasArea returns true if this rectangle has positive width and height. See also [Area].
func HasArea(rect PositionSize) bool { //gd:Rect2i.has_area
	return rect.Size.X > 0 && rect.Size.Y > 0
}

// HasPoint returns true if the rectangle contains the given point. By convention, points
// on the right and bottom edges are not included.
//
// Note: This method is not reliable for Rect2 with a negative size. Use abs first to get
// a valid rectangle.
func HasPoint(rect PositionSize, point Vector2i.XY) bool { //gd:Rect2i.has_point
	if point.X < rect.Position.X {
		return false
	}
	if point.Y < rect.Position.Y {
		return false
	}
	if point.X >= rect.Position.X+rect.Size.X {
		return false
	}
	if point.Y >= rect.Position.Y+rect.Size.Y {
		return false
	}
	return true
}

// Intersection returns the intersection between this rectangle and b. If the rectangles do
// not intersect, returns an empty [PositionSize].
//
// Note: If you only need to know whether two rectangles are overlapping, use [Overlaps].
func Intersection(a, b PositionSize) PositionSize { //gd:Rect2i.intersection
	var new_rect = b
	if !Overlaps(a, b) {
		return PositionSize{}
	}
	new_rect.Position.X = max(b.Position.X, a.Position.X)
	new_rect.Position.Y = max(b.Position.Y, a.Position.Y)
	var (
		rect_end = Vector2i.Add(b.Position, b.Size)
		end      = Vector2i.Add(a.Position, a.Size)
	)
	new_rect.Size.X = min(rect_end.X, end.X) - new_rect.Position.X
	new_rect.Size.Y = min(rect_end.Y, end.Y) - new_rect.Position.Y
	return new_rect
}

// Overlaps returns true if this rectangle overlaps with the b rectangle. The edges of
// both rectangles are excluded.
func Overlaps(a, b PositionSize) bool { //gd:Rect2i.intersects
	const include_borders = false // TODO new function?
	if include_borders {
		if a.Position.X > (b.Position.X + b.Size.X) {
			return false
		}
		if (a.Position.X + a.Size.X) < b.Position.X {
			return false
		}
		if a.Position.Y > (b.Position.Y + b.Size.Y) {
			return false
		}
		if (a.Position.Y + b.Size.Y) < b.Position.Y {
			return false
		}
	} else {
		if a.Position.X >= (b.Position.X + b.Size.X) {
			return false
		}
		if (a.Position.X + a.Size.X) <= b.Position.X {
			return false
		}
		if a.Position.Y >= (b.Position.Y + b.Size.Y) {
			return false
		}
		if (a.Position.Y + a.Size.Y) <= b.Position.Y {
			return false
		}
	}
	return true
}

// Merge returns a [PositionSize] that encloses both this rectangle and b around the edges.
// See also [Inside].
func Merge(a, b PositionSize) PositionSize { //gd:Rect2i.merge
	var new_rect PositionSize
	new_rect.Position.X = min(b.Position.X, a.Position.X)
	new_rect.Position.Y = min(b.Position.Y, a.Position.Y)
	new_rect.Size.X = max(b.Position.X+b.Size.X, a.Position.X+a.Size.X)
	new_rect.Size.Y = max(b.Position.Y+b.Size.Y, a.Position.Y+a.Size.Y)
	new_rect.Size = Vector2i.Sub(new_rect.Size, new_rect.Position) // Make relative again.
	return new_rect
}
